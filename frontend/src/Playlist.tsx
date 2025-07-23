import type { Release } from './client';
import './Playlist.css'

interface PlaylistProps {
	setCurrentTrackUrl: (url: string) => void;
	highLightedTrack: string | null;
	release: Release;
}
const Playlist: React.FC<PlaylistProps> = ({ release: release, setCurrentTrackUrl: setCurrentTrackUrl, highLightedTrack: highLightedTrack }) => {
	return (
		<div className='playlist'>
			{!release || !release.tracks || release.tracks.length === 0 ? (
				<p>No songs found.</p>
			) : (
				<div style={{ margin: '10px' }}>
					<div className='album'>
						<img src={release.cover} width="200" height="200" className='rad-shadow albumCoverImg' />
						<h1 style={{ marginLeft: '60px', marginTop: '60px', color: 'var(--text1)' }}>{release?.name}</h1>
					</div>
					<table style={{ width: '100%', borderCollapse: 'collapse', borderSpacing: '0 10px', marginTop: '30px' }}>
						<thead style={{ borderBottom: '2px solid var(--surface2)' }}>
							<tr style={{ color: 'var(--text2)' }}>
								<th style={{ textAlign: 'left' }}>Song</th>
								<th style={{ textAlign: 'right' }}>Duration</th>
							</tr>
						</thead>
						<tbody>
							{release.tracks.map(track => (
								<tr onClick={() => setCurrentTrackUrl(track.file)} style={{ cursor: 'pointer' }}>
									<td key={track.file} style={{ textAlign: 'left', color: 'var(--text1)' }}>
										{highLightedTrack != '' && track.file == highLightedTrack ? (
											<strong className='current-song'>{track.name}</strong>
										) : (
											<strong>{track.name}</strong>
										)}
									</td>
									<td style={{ textAlign: 'right', color: 'var(--text2)' }}>{formatTime(track.duration)}</td>
								</tr>
							))}
						</tbody>
					</table>
				</div>
			)}
		</div>
	);
};

function formatTime(time: number) {
	const minutes = Math.floor(time / 60);
	const seconds = Math.floor(time % 60).toString().padStart(2, '0');
	return `${minutes}:${seconds}`;
}


export default Playlist;

