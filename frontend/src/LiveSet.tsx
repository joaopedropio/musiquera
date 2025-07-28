import type { Release, Segment } from './client';
import './LiveSet.css'

interface LiveSetProps {
	setCurrentTrackUrl: (currentTrackUrl: string) => void;
	setCurrentSegment: (currentSegment: Segment) => void;
	setCurrentReleaseType: (currentReleaseType: string) => void;
	highLightedTrack: string | null;
	release: Release;
}
const LiveSet: React.FC<LiveSetProps> = ({ release: release, setCurrentTrackUrl: setCurrentTrackUrl, setCurrentSegment: setCurrentSegment, setCurrentReleaseType: setCurrentReleaseType, highLightedTrack: highLightedTrack }) => {
	return (
		<div className='liveSet'>
			{!release || !release.tracks || release.tracks.length === 0 ? (
				<p>No songs found.</p>
			) : (
				<div style={{ margin: '10px' }}>
					<div className='liveSetHeader'>
						<img src={release.cover} width="200" height="200" className='rad-shadow liveSetCoverImg' />
						<h1 style={{ marginLeft: '60px', marginTop: '60px', color: 'var(--text1)' }}>{release?.name}</h1>
					</div>
					<table style={{ width: '100%', borderCollapse: 'collapse', borderSpacing: '0 10px', marginTop: '30px' }}>
						<thead style={{ borderBottom: '2px solid var(--surface2)' }}>
							<tr style={{ color: 'var(--text2)' }}>
								<th style={{ textAlign: 'left' }}>Segment</th>
								<th style={{ textAlign: 'right' }}>Position</th>
							</tr>
						</thead>
						<tbody>
							{release.tracks[0].segments.map(segment => (
								<tr onClick={() => {setCurrentSegment(segment);setCurrentReleaseType('liveSet');setCurrentTrackUrl(release.tracks[0].file)}} style={{ cursor: 'pointer' }}>
									<td key={segment.name} style={{ textAlign: 'left', color: 'var(--text1)' }}>
										{highLightedTrack != '' && segment.name == highLightedTrack ? (
											<strong className='current-segment'>{segment.name}</strong>
										) : (
											<strong>{segment.name}</strong>
										)}
									</td>
									<td style={{ textAlign: 'right', color: 'var(--text2)' }}>{formatTime(segment.position)}</td>
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


export default LiveSet;

