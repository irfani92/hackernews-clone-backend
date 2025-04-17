// lib/hn.ts

const BASE_URL = 'http://localhost:8080';

export async function getStories(type: string, page = 1, pageSize = 10) {
    const allowedTypes = ['topstories', 'newstories', 'askstories', 'showstories', 'jobstories'];
    const endpoint = allowedTypes.includes(type) ? type : 'topstories';

    const res = await fetch(`${BASE_URL}/${endpoint}?page=${page}&pageSize=${pageSize}`, {
        cache: 'no-store',
    });

    if (!res.ok) throw new Error('Failed to fetch stories');
    return res.json();
}