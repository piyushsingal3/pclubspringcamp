
import React, { useEffect, useState } from 'react';

function UserData() {
  const [userData, setUserData] = useState([]);

  useEffect(() => {
    // fetch recent actions from Go server
    fetch('http://localhost:9000/userData')
      .then(response => response.json())
      .then(data => setUserData(data))
      .catch(error => console.error('Error fetching user data:', error));
  }, []); 

  return (
    <div>
      <h1>Users Data</h1>
      {userData.map(entry => (
        <div key={entry.handle}>
          <div>
            <h3>Handle: {entry.handle}</h3>
            <p>Username: {entry.userName}</p>
            <p>Email: {entry.email}</p>
          
            {entry.subscribedBlogs && entry.subscribedBlogs.length > 0 && (
              <div>
                <p>Subscribed Blogs:</p>
                <ul>
                  {entry.subscribedBlogs.map((blog, index) => (
                    <li key={index}>{blog}</li>
                  ))}
                </ul>
              </div>
            )}
          </div>
        </div>
      ))}
    </div>
  );
}

export default UserData;
