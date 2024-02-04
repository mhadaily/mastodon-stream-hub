import React, { useState, useEffect, useRef } from 'react';

// Todo(mhadaily): this can be replaced by TS generated from Protobuf
interface Message {
	id: string;
	url: string;
	account: {
		display_name: string;
	};
	content: string;
}

const RealtimeFeed: React.FC = () => {
	const [messages, setMessages] = useState<Message[]>([]);
	const topOfListRef = useRef(null);

	useEffect(() => {
		// Access the environment variable for the WebSocket endpoint
		const wsEndpoint =
			import.meta.env.REACT_APP_WS_ENDPOINT || 'ws://localhost:8080/ws';
		const ws = new WebSocket(wsEndpoint);
		ws.binaryType = 'arraybuffer';

		ws.onopen = () => {
			console.log('WebSocket Connected');
		};

		ws.onmessage = (event) => {
			if (typeof event.data === 'string') {
				console.log('Received text message:', event.data);
			} else {
				const arrayBuffer = event.data;
				const decoder = new TextDecoder('utf-8');
				const text = decoder.decode(arrayBuffer);
				const message: Message = JSON.parse(text);
				setMessages((prevMessages) => prevMessages.concat(message));
			}
		};

		ws.onclose = () => {
			console.log('WebSocket Disconnected');
		};

		if (topOfListRef.current) {
			topOfListRef.current.scrollIntoView({ behavior: 'smooth' });
		}

		return () => {
			ws.close();
		};
	}, [messages]);

	return (
		<div>
			<h2 className="text-3xl font-bold text-center">Real-time Feed</h2>
			<ul>
				{messages.map((msg, index) => (
					<li key={index}>
						<div className="px-10 my-4 py-6 rounded shadow-xl bg-white w-4/5 mx-auto">
							<div className="flex justify-between items-center">
								<span className="font-light text-gray-600">{}</span>
								<a
									className="px-2 py-1 bg-blue-600 text-gray-100 font-bold rounded hover:bg-gray-500"
									href="#">
									{msg.account.display_name}
								</a>
							</div>
							<div className="mt-2">
								<a
									className="text-2xl text-gray-700 font-bold hover:text-gray-600"
									href="#">
									{msg.id}
								</a>
								<p
									className="mt-2 text-gray-600"
									dangerouslySetInnerHTML={{ __html: msg.content }}></p>
							</div>
							<div className="flex justify-between items-center mt-4">
								<a className="text-blue-600 hover:underline" href={msg.url}>
									Visit Article
								</a>
							</div>
						</div>
					</li>
				))}
			</ul>
			<div ref={topOfListRef}></div>
		</div>
	);
};

export default RealtimeFeed;
