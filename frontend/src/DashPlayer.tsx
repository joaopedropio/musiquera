import React, { useRef, useState, useEffect } from 'react';
import { MediaPlayer } from 'dashjs'

interface DashPlayerProps {
	src: string;
	autoplay?: boolean;
}

const DashPlayer: React.FC<DashPlayerProps> = ({ src, autoplay }) => {
	const audioRef = useRef<HTMLAudioElement | null>(null);
	const [isPlaying, setIsPlaying] = useState(false);
	const [volume, setVolume] = useState(1); // Volume range: 0 to 1
	const [currentTime, setCurrentTime] = useState(0);
	const [duration, setDuration] = useState(0);

	const handlePlay = () => {
		if (audioRef.current) {
			audioRef.current.play();
			setIsPlaying(true);
		}
	};

	const handlePause = () => {
		if (audioRef.current) {
			audioRef.current.pause();
			setIsPlaying(false);
		}
	};

	const handleVolumeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const newVolume = parseFloat(e.target.value);
		if (audioRef.current) {
			audioRef.current.volume = newVolume;
			setVolume(newVolume);
		}
	};

	const handleSeek = (e: React.ChangeEvent<HTMLInputElement>) => {
		const newTime = parseFloat(e.target.value);
		if (audioRef.current) {
			audioRef.current.currentTime = newTime;
		}
		setCurrentTime(newTime);
	};

	useEffect(() => {
		if (audioRef.current && src) {
			const audio = audioRef.current;
			const player = MediaPlayer().create();
			player.initialize(audioRef.current, src, autoplay);
			const updateTime = () => setCurrentTime(audio.currentTime);
			const setAudioDuration = () => setDuration(audio.duration);

			audio.addEventListener('timeupdate', updateTime);
			audio.addEventListener('loadedmetadata', setAudioDuration);

			return () => {
				audio.removeEventListener('timeupdate', updateTime);
				audio.removeEventListener('loadedmetadata', setAudioDuration);
				player.reset();
			};
		}
	}, [src, autoplay]);

	return (
		<div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
			<audio ref={audioRef} preload="auto" />

			<div style={{ marginBottom: '10px' }}>
				<button onClick={handlePlay} disabled={isPlaying}>Play</button>
				<button onClick={handlePause} disabled={!isPlaying}>Pause</button>
			</div>

			<div>
				<label>
					Volume:
					<input
						type="range"
						min="0"
						max="1"
						step="0.01"
						value={volume}
						onChange={handleVolumeChange}
					/>
				</label>
			</div>
			<div>
				<label>
					Progress:
					<input
						type="range"
						min="0"
						max={duration || 0}
						step="0.1"
						value={currentTime}
						onChange={handleSeek}
						style={{ width: '100%' }}
					/>
				</label>
				<div style={{ fontSize: '0.9em', marginTop: '4px' }}>
					{formatTime(currentTime)} / {formatTime(duration)}
				</div>
			</div>
		</div>
	);
};

function formatTime(time: number) {
	const minutes = Math.floor(time / 60);
	const seconds = Math.floor(time % 60).toString().padStart(2, '0');
	return `${minutes}:${seconds}`;
}


export default DashPlayer;

