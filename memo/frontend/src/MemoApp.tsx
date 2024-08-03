import React, { useState, useEffect, ChangeEvent } from 'react';

interface Memo {
    id: number;
    content: string;
}

const MemoApp: React.FC = () => {
    const [memos, setMemos] = useState<Memo[]>([]);
    const [content, setContent] = useState<string>('');

    useEffect(() => {
        fetchMemos();
    }, []);

    const fetchMemos = async () => {
        const response = await fetch('http://34.152.37.120:8080/memos');
        const data = await response.json();
        setMemos(data);
    };

    const addMemo = async () => {
        await fetch('http://34.152.37.120:8080/memo', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ content }),
        });
        setContent('');
        fetchMemos();
    };

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        setContent(e.target.value);
    };

    return (
        <div>
            <h1>Memo App</h1>
            <input
                type="text"
                value={content}
                onChange={handleChange}
            />
            <button onClick={addMemo}>Add Memo</button>
            <ul>
                {memos.map((memo) => (
                    <li key={memo.id}>{memo.content}</li>
                ))}
            </ul>
        </div>
    );
};

export default MemoApp;