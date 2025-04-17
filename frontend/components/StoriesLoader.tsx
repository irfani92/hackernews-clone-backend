// components/StoriesLoader.tsx
'use client';

import { useEffect, useState } from 'react';
import StoryItem from './StoryItem';

type Story = {
    id: number;
    title: string;
    url?: string;
    by: string;
    score: number;
};

type Props = {
    type: string;
    page: number;
    pageSize: number;
};

export default function StoriesLoader({ type, page, pageSize }: Props) {
    const [stories, setStories] = useState<Story[] | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        setStories(null); // reset when type/page changes
        setLoading(true);

        fetch(`http://localhost:8080/${type}stories?page=${page}&pageSize=${pageSize}`)
            .then(res => res.json())
            .then(data => {
                setStories(data);
                setLoading(false);
            });
    }, [type, page]);

    if (loading || !stories) {
        return (
            <div className="flex justify-center items-center py-10">
                <div className="w-10 h-10 border-4 border-orange-500 border-t-transparent rounded-full animate-spin"></div>
            </div>
        );
    }

    return (
        <>
            {stories.map((story, i) => (
                <StoryItem key={story.id} story={story} index={(page - 1) * pageSize + i} />
            ))}
        </>
    );
}
