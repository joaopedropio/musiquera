import React, { useRef, useState, useEffect } from 'react';
import { MediaPlayer } from 'dashjs'
import './DashPlayer.css'
import type { Segment } from './client';

interface DashPlayerProps {
	src: string;
	currentSegment: Segment
	autoplay?: boolean;
	onNextSegment: () => void;
	onPreviousSegment: () => void;
}


const DashPlayerLiveSet: React.FC<DashPlayerProps> = ({ src, currentSegment: currentSegment, onNextSegment: onNextSegment, autoplay, onPreviousSegment: onPreviousSegment }) => {
	const audioRef = useRef<HTMLAudioElement | null>(null);
	const [isPlaying, setIsPlaying] = useState(false);
	const [volume, setVolume] = useState(1); // Volume range: 0 to 1
	const [currentTime, setCurrentTime] = useState(0);
	const [duration, setDuration] = useState(0);

	const onPlay = () => {
		if (audioRef.current) {
			setIsPlaying(true);
		}
	}

	const onPause = () => {
		if (audioRef.current) {
			setIsPlaying(false);
		}
	}

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

			audio.currentTime = currentSegment.position
			audio.addEventListener('timeupdate', updateTime);
			audio.addEventListener('loadedmetadata', setAudioDuration);

			audio.addEventListener('play', onPlay);
			audio.addEventListener('pause', onPause);

			return () => {
				audio.removeEventListener('timeupdate', updateTime);
				audio.removeEventListener('loadedmetadata', setAudioDuration);
				audio.removeEventListener('play', onPlay);
				audio.removeEventListener('pause', onPause);
				player.reset();
			};
		}
	}, [src, autoplay]);

	useEffect(() => {
		if (audioRef.current) {
			audioRef.current.currentTime = currentSegment.position
		}
	}, [currentSegment])

	return (
		<div className='dashPlayer'>
			<audio ref={audioRef} preload="auto" />
			<div className='player-controls'>
				<div>
					<button onClick={() => onPreviousSegment()} className="icon-button" title="Previous">
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
							<path d="M6 12l10 7V5zM4 5h2v14H4z" />
						</svg>
					</button>
					{!isPlaying ? (
						<button onClick={handlePlay} className="icon-button" title="Play">
							<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
								<path d="M8 5v14l11-7z" />
							</svg>
						</button>
					) : (
						<button onClick={handlePause} className="icon-button" title="Pause">
							<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
								<path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z" />
							</svg>
						</button>
					)}
					<button onClick={() => onNextSegment()} className="icon-button" title="Next">
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
							<path d="M18 12L8 5v14zM20 5h-2v14h2z" />
						</svg>
					</button>
				</div>
				<div>
					<label>
						<strong>{formatTime(currentTime)}</strong>
						<input
							className='progressBar'
							type="range"
							min="0"
							max={duration || 0}
							step="0.1"
							value={currentTime}
							onChange={handleSeek}
							style={{ width: '60%' }}
						/>
						<strong>{formatTime(duration)}</strong>
					</label>
				</div>
			</div>
			<div className='playerOptions'>
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

		</div>
	);
};

function formatTime(time: number) {
	const minutes = Math.floor(time / 60);
	const seconds = Math.floor(time % 60).toString().padStart(2, '0');
	return `${minutes}:${seconds}`;
}


export default DashPlayerLiveSet;

