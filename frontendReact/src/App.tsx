import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Home } from './pages/home';
import { Login } from './pages/login';
import { Signup } from './pages/signup';

function App() {
  const [islogin, setLogin] = useState(false)

  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path='/login' element={<Login setLogIn={setLogin} isLogin={islogin}/>} />
          <Route path='/signup' element={<Signup />} />
          <Route path='/home' element={<Home />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
