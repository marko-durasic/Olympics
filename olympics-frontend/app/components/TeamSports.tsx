"use client";

import { useEffect, useState } from 'react';

interface TeamSport {
  id: number;
  country: string;
  sport: string;
  gold: number;
  silver: number;
  bronze: number;
}

const TeamSports = () => {
  const [teamSports, setTeamSports] = useState<TeamSport[]>([]);

  useEffect(() => {
    fetch('http://localhost:8080/api/team-sports')
      .then(response => response.json())
      .then(data => {
        console.log('Fetched team sports data:', data);
        setTeamSports(data);
      })
      .catch(error => console.error('Error fetching team sports:', error));
  }, []);

  return (
    <div className="p-4 bg-white shadow-md rounded-lg w-full max-w-2xl">
      <h2 className="text-2xl font-bold mb-4 text-gray-800">Team Sports</h2>
      <ul className="list-disc pl-5">
        {teamSports.length > 0 ? (
          teamSports.map(sport => (
            <li key={sport.id} className="mb-2">
              <span className="font-semibold text-blue-600">{sport.country}</span> - 
              <span className="text-green-600">{sport.sport}</span>: 
              <span className="text-yellow-500"> {sport.gold} Gold</span>, 
              <span className="text-gray-500"> {sport.silver} Silver</span>, 
              <span className="text-orange-500"> {sport.bronze} Bronze</span>
            </li>
          ))
        ) : (
          <li>Loading...</li>
        )}
      </ul>
    </div>
  );
};

export default TeamSports;
