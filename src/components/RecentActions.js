import React, { useEffect, useState } from 'react';

function RecentActions() {
  const [blogData, setBlogData] = useState([]);

  useEffect(() => {
    // fetch recent actions from Go server
    fetch('http://localhost:9000/Activity/recentActions')
      .then(response => response.json())
      .then(data => {console.log(data);setBlogData(data)})
      .catch(error => console.error('Error fetching recent actions:', error));
  }, []); 

  return (
    <div>
      <h1>Recent Blog Entries</h1>
      {blogData.map(entry => (
        <div key={entry.blogEntry.id}>
          
          <h3>Title: {entry.blogEntry.title}</h3>
          <p>TimeSecond:{entry.timeSeconds}</p>
          <p>ID: {entry.blogEntry.id}</p>
          <p>Original Locale: {entry.blogEntry.originalLocale}</p>
          <p>Creation Time: {entry.blogEntry.creationTimeSeconds}</p>
          
        </div>
      ))}
    </div>
  );
}

export default RecentActions;