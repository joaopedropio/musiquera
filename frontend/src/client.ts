import axios from 'axios'

//const ReleaseTypeAlbum = 'album'
//const ReleaseTypeLiveSet = 'liveSet'

type ReleaseType = string

export type Segment = {
	name: string;
	position: number;
}

export type Track = {
	name: string;
	file: string;
	duration: number;
	segments: Segment[];
}

export type Release = {
	name: string;
	cover: string;
	type: ReleaseType;
	releaseDate: string;
	tracks: Track[];
}

export type Artist = {
	name: string;
	profileCoverPath: string;
}

export class Client {
	constructor() {

	}

	async getAllArtists(): Promise<Artist[]> {
		const response = await axios.get<Artist[]>('/api/artist/')
		return response.data
	}

	async getReleasesByArtist(artistName: string): Promise<Release[]> {
		const response = await axios.get<Release[]>('/api/release/byArtist/' + artistName)
		return response.data
	}

	async getMostRecentRelease(): Promise<Release> {
		const resp = await axios.get<Release>('/api/release/mostRecent')
		return resp.data
	}
}
