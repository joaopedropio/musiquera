import { useState } from 'react'
import './App.css'
import './DashPlayer'
import DashPlayer from './DashPlayer'
import Playlist from './Playlist'
import ArtistComponent from './Artist'

function App() {
	const [currentSongUrl, setCurrentSongUrl] = useState<string | null>(null);
	return (
		<>
			<div>
				<ArtistComponent />
				<Playlist setCurrentSongUrl={setCurrentSongUrl} />
			</div>
			<div>
				{currentSongUrl && <DashPlayer src={currentSongUrl} autoplay />}
			</div>
		</>
	)
}

export default App
