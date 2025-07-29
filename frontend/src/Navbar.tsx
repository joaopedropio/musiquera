import { useNavigate } from 'react-router-dom'
import './Navbar.css'
const Navbar = () => {
	const navigate = useNavigate()

	const handleLogout = async () => {
		await fetch('/logout', {
			method: 'POST',
			credentials: 'include',
		})

		navigate('/loginPage')
	}

	return (
		<div className="nav rad-shadow">
			<nav className="navbar">
				<div>Musiquera</div>
				<div className='logoutButton'>
					<button onClick={handleLogout}>Logout</button>
				</div>
			</nav>
		</div>
	)
}

export default Navbar

