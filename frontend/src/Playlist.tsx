import type { Album } from './client';
import './Playlist.css'

interface PlaylistProps {
	setCurrentSongUrl: (url: string) => void;
	album: Album;
}
const Playlist: React.FC<PlaylistProps> = ({ album, setCurrentSongUrl }) => {
	return (
		<div className='playlist'>
			{!album || !album.songs || album.songs.length === 0 ? (
				<p>No songs found.</p>
			) : (
				<div>
					<h2>{album?.name}</h2>
					<ul>
						{album.songs.map(song => (
							<li key={song.file} onClick={() => setCurrentSongUrl(song.file)} style={{ cursor: 'pointer' }}>
								<strong>{song.name}</strong>
							</li>
						))}
					</ul>
				</div>
			)}
		</div>
	);
};

export default Playlist;

