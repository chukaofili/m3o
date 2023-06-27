import { NextApiRequest, NextApiResponse } from 'next'
import call from '../../../../lib/micro'
import TokenFromReq from '../../../../lib/token'

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse
) {
    const {
        query: { group_id },
    } = req

    // ignore any OPTIONS requests
    if (!['GET', 'POST']?.includes(req.method)) {
        res.status(200)
        return
    }

    // get the token from cookies
    const token = TokenFromReq(req)
    if (!token) {
        res.status(401).json({ error: 'No token cookie set' })
        return
    }

    // authenticate the request
    let user: any
    try {
        const rsp = await call('/users/validate', { token })
        user = rsp.user
    } catch ({ error, code }) {
        const statusCode = code === 400 ? 401 : code
        res.status(statusCode).json({ error })
        return
    }

    // load the groups the user is a part of
    let group: any
    try {
        const rsp = await call('/groups/List', { member_id: user.id })
        group = rsp.groups?.find((g) => g.id === group_id)
    } catch ({ error, code }) {
        console.error(`Error loading groups: ${error}, code: ${code}`)
        res.status(500).json({ error: 'Error loading groups' })
        return
    }
    if (group) {
        res.status(200).json({})
        return
    }

    // load the group

    try {
      group = (await call('/groups/Read', { ids: [group_id] })).groups[group_id as string]
    } catch ({ error, code }) {
      console.error(`Error reading group: ${error}, code: ${code}`)
      res.status(500).json({ error: 'Error reading group' })
      return
    }

    // add the user as a member of the group
    try {
      await call('/groups/AddMember', {
        group_id: group_id,
        member_id: user.id,
      })
    } catch ({ error, code }) {
      console.error(`Error adding member ${user.id} to group ${group_id}: ${error}`)
      res.status(500).json({ error: 'Error accepting invitation' })
      return
    }

    // publish the message to the users in the group
    try {
      group.member_ids.forEach(async (id: string) => {
        await call('/streams/Publish', {
          topic: id,
          message: JSON.stringify({
            type: 'group.user.joined',
            group_id: group.id,
            payload: user,
          }),
        })
      })
      res.status(200).json({})
    } catch ({ error, code }) {
      console.error(`Error publishing to stream: ${error}, code: ${code}`)
      res.status(500).json({ error: 'Error publishing to stream' })
    }
}
