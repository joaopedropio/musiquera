import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import logo from '../src/assets/favicon/android-chrome-512x512.png'
import './Login.css'

const Login: React.FC = () => {
	const [username, setUsername] = useState('')
	const [password, setPassword] = useState('')
	const [error, setError] = useState('')
	const navigate = useNavigate()

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault()

		try {
			const res = await fetch('/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				credentials: 'include', // sends & receives cookies
				body: JSON.stringify({ "username": username, "password": password }),
			})

			if (res.ok) {
				navigate('/') // redirect to app
			} else {
				const text = await res.text()
				setError(`Login failed: ${text}`)
			}
		} catch (err) {
			setError('Network error. Please try again.')
		}
	}

	return (
		<div style={{ maxWidth: 600, margin: 'auto', paddingTop: '2rem', display: 'flex', flexDirection: 'column', justifyContent: 'center' }}>
			<div className='loginLogoImg logoRadShadow'>
				<img src={logo} style={{ maxWidth: '200px' }} />
			</div>
			<div>
				<form onSubmit={handleSubmit}>
					<input
						type="text"
						placeholder="Username"
						value={username}
						onChange={e => setUsername(e.target.value)}
						required
						style={{ width: '100%', padding: '0.5rem', marginBottom: '1rem' }}
					/>
					<input
						type="password"
						placeholder="Password"
						value={password}
						onChange={e => setPassword(e.target.value)}
						required
						style={{ width: '100%', padding: '0.5rem', marginBottom: '1rem' }}
					/>
					<button type="submit" style={{ width: '100%', padding: '0.5rem' }}>
						Log In
					</button>
				</form>
			</div>
			{error && <p style={{ color: 'red' }}>{error}</p>}
		</div>
	)
}

export default Login

