// app/page.tsx
import StoriesLoader from '../components/StoriesLoader';
import Link from 'next/link';

type SearchParams = {
  searchParams?: {
    page?: string;
    type?: string;
  };
};

const storyTypes = {
  top: 'Top',
  new: 'New',
  ask: 'Ask',
  show: 'Show',
  job: 'Job',
};

export default function HomePage({ searchParams }: SearchParams) {
  const page = parseInt(searchParams?.page || '1', 10);
  const type = searchParams?.type || 'top';
  const pageSize = 10;

  return (
    <main className="w-full min-h-screen bg-gray-100 text-gray-900 flex flex-col items-center">
      <div className="w-full max-w-4xl px-4 md:px-6 py-6">
        <h1 className="text-3xl font-bold mb-6 text-orange-600 text-center">
          Hacker News Clone
        </h1>

        <nav className="mb-6 flex flex-wrap justify-center gap-4 text-sm">
          {Object.entries(storyTypes).map(([key, label]) => (
            <Link
              key={key}
              href={`/?type=${key}&page=1`}
              className={`px-3 py-1 border-b-2 transition ${type === key
                ? 'border-orange-500 font-semibold text-orange-600'
                : 'border-transparent text-gray-600 hover:border-gray-400 hover:text-gray-800'
                }`}
            >
              {label}
            </Link>
          ))}
        </nav>

        <StoriesLoader type={type} page={page} pageSize={pageSize} />

        <div className="mt-10 flex justify-between text-sm">
          {page > 1 && (
            <Link href={`/?type=${type}&page=${page - 1}`} className="text-blue-600 hover:underline">
              ← Prev
            </Link>
          )}
          <Link href={`/?type=${type}&page=${page + 1}`} className="text-blue-600 hover:underline ml-auto">
            Next →
          </Link>
        </div>
      </div>
    </main>
  );
}
