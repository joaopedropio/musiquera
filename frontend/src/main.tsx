import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Login from './Login.tsx'
import './index.css'
import App from './App.tsx'
import RequireAuth from './RequireAuth.tsx'
import Dashboard from './dashboard/Dashboard.tsx'

async function enableMocking() {
	if (process.env.NODE_ENV !== 'development') {
		return;
	}

	const { worker } = await import('../mocks/browser')
	return worker.start()
}

enableMocking().then(() => {
	createRoot(document.getElementById('root')!).render(
		<StrictMode>
			<Router>
				<Routes>
					<Route path="/" element={
						<RequireAuth>
							<App />
						</RequireAuth>
					} />
					<Route path="/loginPage" element={<Login />} />
					<Route path="/dashboard" element={<Dashboard />} />
				</Routes>
			</Router>
		</StrictMode>,
	)
})

