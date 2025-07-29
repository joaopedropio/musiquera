import { useEffect, useState, type ReactElement } from 'react'
import { Navigate } from 'react-router-dom'
import { isAuthenticated } from './auth'

const RequireAuth = ({ children }: { children: ReactElement }) => {
	const [loading, setLoading] = useState(true)
	const [authed, setAuthed] = useState(false)

	useEffect(() => {
		isAuthenticated().then(result => {
			setAuthed(result)
			setLoading(false)
		})
	}, [])

	if (loading) return <div>Loading...</div>

	return authed ? children : <Navigate to="/loginPage" />
}

export default RequireAuth

