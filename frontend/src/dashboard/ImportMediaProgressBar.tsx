import React, { useEffect, useState } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';

const ProgressMessages: React.FC = () => {
	const webSocketURL = getWebSocketUrl()
	const { lastJsonMessage, readyState } = useWebSocket(webSocketURL)
	const [jobs, setJobs] = useState<Map<string, Job>>(new Map())

	const updateJob = (newJob: Job) => {
		setJobs((prevJobs) => {
			// Create a copy of the Map to avoid direct mutation
			const updatedJobs = new Map(prevJobs);
			updatedJobs.set(newJob.job_id, newJob);  // Override job by job_id
			return updatedJobs;
		});
	};

	const convertToJob = (message: any): Job | null => {
		if (!message || !message.job_id || !message.progress) {
			// In case the message doesn't match the expected structure, return null
			return null;
		}

		// Assuming message has the necessary fields, we map them to the Job type
		const job: Job = {
			job_id: message.job_id,
			started: message.started ?? false, // Default to false if not provided
			finished: message.finished ?? false, // Default to false if not provided
			progress: {
				percentage: message.progress.percentage || "0%", // Default to "0%" if missing
				size: message.progress.size || "0MB", // Default to "0MB" if missing
				speed: message.progress.speed || "0MB/s", // Default to "0MB/s" if missing
				eta: message.progress.eta || "N/A", // Default to "N/A" if missing
			},
		};

		return job;
	};

	useEffect(() => {
		if (lastJsonMessage) {
			console.log('Received message:', JSON.stringify(lastJsonMessage))
			const job = convertToJob(lastJsonMessage)
			if (job) {
				updateJob(job)
			}
		}
	}, [lastJsonMessage])

	return (
		<div>
			<h2>WebSocket Status: {readyState === ReadyState.OPEN ? 'Connected' : 'Disconnected'}</h2>
			<div>
				{/*
				<h3>Received Message:</h3>
				<p>{lastJsonMessage ? JSON.stringify(lastJsonMessage) : 'No messages yet'}</p>
			*/}
				{Array.from(jobs.values()).map((job) => (
					<div key={job.job_id}>
						<h4>Job ID: {job.job_id}</h4>
						<p>Status: {job.started ? (job.finished ? "Finished" : "Started") : "Not Started"}</p>
						<p>Progress: {job.progress.percentage} at {job.progress.speed} speed</p>
					</div>
				))}
			</div>
		</div>
	)
}

type Progress = {
	percentage: string,
	size: string,
	speed: string,
	eta: string,
}

type Job = {
	job_id: string,
	started: boolean,
	finished: boolean,
	progress: Progress,
}

const getWebSocketUrl = (): string => {
	if (process.env.NODE_ENV == 'development') {
		return "ws://localhost:8080/progress"
	}
	const baseUrl = window.location.origin
	return `${baseUrl.replace(/^http/, 'ws')}/progress`
}

export default ProgressMessages;

