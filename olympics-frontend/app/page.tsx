"use client";

import TeamSports from './components/TeamSports';

export default function Home() {
    return (
        <main className="flex min-h-screen flex-col items-center justify-center p-4 bg-gray-100">
            <h1 className="text-4xl font-bold mb-4 text-blue-600">Olympics Data Viewer</h1>
            <TeamSports />
        </main>
    );
}
