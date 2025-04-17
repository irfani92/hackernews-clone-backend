// app/loading.tsx
export default function Loading() {
    return (
        <div className="flex justify-center items-center min-h-screen">
            <div className="flex flex-col items-center">
                <div className="w-12 h-12 border-4 border-orange-500 border-t-transparent rounded-full animate-spin mb-4"></div>
                <p className="text-gray-600 text-sm">Loading stories...</p>
            </div>
        </div>
    );
}
