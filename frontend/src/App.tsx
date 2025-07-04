import { useState } from 'react'
import './App.css'
import './DashPlayer'
import DashPlayer from './DashPlayer'
import Playlist from './Playlist'

function App() {
	const [currentSongUrl, setCurrentSongUrl] = useState<string | null>(null);
	return (
		<>
			<Playlist setCurrentSongUrl={setCurrentSongUrl}/>
			{ currentSongUrl && <DashPlayer src={currentSongUrl} autoplay />}
		</>
	)
}

export default App
