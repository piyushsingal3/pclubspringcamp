import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';



import './App.css';

function AfterSignin() {
  return (
    
       
    
    <div className="App"> 
          <nav>
          <ul>
            <li>
              <Link to="/">Sign in </Link>
            </li>
            <li>
              <Link to="/about">About</Link>
            </li>
            <li>
              <Link to="/services">Services</Link>
            </li>
            <li>
              <Link to="/Activity/recentActions">RecentActions</Link>
            </li>
            <li>
              <Link to="/userData">User</Link>
            </li>
          </ul>
        </nav> 
      </div>


  );
}

export default AfterSignin;
