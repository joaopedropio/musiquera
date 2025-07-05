import type { Artist } from './client'

interface Props {
	artists: Artist[]
	setCurrentArtist: (artist: string) => void
}

const ArtistComponent: React.FC<Props> = ({artists, setCurrentArtist}) => {
//	const [allArtists, setAllArtists] = useState<Artist[]>([])
//	const [loading, setLoading] = useState<boolean>(true)
//	const [error, setError] = useState<string | null>(null)

//	useEffect(() => {
//		const fetchArtists = async () => {
//			try {
//				const client = new Client()
//				const artists = await client.getAllArtists()
//				setAllArtists(artists)
//
//			} catch (err) {
//				setError((err as Error).message)
//			} finally {
//				setLoading(false)
//			}
//
//		}
//		fetchArtists()
//	}, [])
//
//	if (loading) return <div>Loading Artist...</div>
//	if (error) return <div>Error: {error}</div>

	return (
		<div>
			{artists.length === 0 ? (
				<p> No artists found. </p>
			) : (
				<ul>
					{artists.map(artist => (
						<li key={artist.name} onClick={() => setCurrentArtist(artist.name)} style={{cursor: 'pointer'}}>
							<strong>{artist.name}</strong>
						</li>
					))}
				</ul>
			)}

		</div>
	)
}

export default ArtistComponent
