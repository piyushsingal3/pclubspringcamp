import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

import About from './components/About';
import Services from './components/Services';
import UserData from './components/User';
import './App.css';
import RecentActions from './components/RecentActions';
import SignInSide from './components/SignInpg';
import AfterSignin from './components/AfterSignin';
import SignUp from './components/Signup';
function App() {
  return (
    <Router>
        
      <div className="App"> 
     
      


        <Routes>
          <Route path="/" element={<SignInSide />} />
          <Route path="/about" element={<About />} />
          <Route path="/services" element={<Services />} />
          <Route path="//Activity/recentActions" element={<RecentActions />} />
          <Route path="/aftersignin" element={<AfterSignin />}/>
          <Route path="/userData" element={<UserData />} />
          <Route path="/Signup" element={<SignUp />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
