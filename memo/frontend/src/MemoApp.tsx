import React, { useState, useEffect, ChangeEvent } from 'react';

interface Memo {
    ID: number;
    content: string;
}

const MemoApp: React.FC = () => {
    const [memos, setMemos] = useState<Memo[]>([]);
    const [content, setContent] = useState<string>('');

    useEffect(() => {
        fetchMemos();
    }, []);

    const fetchMemos = async () => {
        try {
            const response = await fetch('http://34.152.37.120:8080/memos');
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            setMemos(data);
            console.log(data);
        } catch (error) {
            console.error('Error fetching memos:', error);
        }
    };

    const addMemo = async () => {
        try {
            const response = await fetch('http://34.152.37.120:8080/memo', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ content }),
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            setContent('');
            fetchMemos();
        } catch (error) {
            console.error('Error adding memo:', error);
        }
    };

    const deleteMemo = async (ID: number) => {
        try {
            const response = await fetch(`http://34.152.37.120:8080/memo/${ID}`, {
                method: 'DELETE',
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            fetchMemos();
        } catch (error) {
            console.error('Error deleting memo:', error);
        }
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
                    <li key={memo.ID}>
                        {memo.content}
                        {memo.ID}
                        <button onClick={() => deleteMemo(memo.ID)}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default MemoApp;