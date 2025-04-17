// components/StoryItem.tsx
type StoryProps = {
    story: {
        id: number;
        title: string;
        url?: string;
        by: string;
        score: number;
    };
    index: number;
};

export default function StoryItem({ story, index }: StoryProps) {
    return (
        <div className="border-b border-gray-200 py-2">
            <div>
                <span className="text-gray-500">{index + 1}. </span>
                <a href={story.url} target="_blank" rel="noopener noreferrer" className="font-semibold text-blue-700 hover:underline">
                    {story.title}
                </a>
            </div>
            <div className="text-sm text-gray-500">
                {story.score} points by {story.by}
            </div>
        </div>
    );
}
