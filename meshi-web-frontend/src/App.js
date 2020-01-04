import React from 'react';
import MyAppBar from './MyAppBar.js';
import Contents from './Contents.js';
import './App.css';

class App extends React.Component {
  render() {
    return (
      <div>
        <MyAppBar />
        <Contents meshi='shokujin' />
      </div>
    );
  }
}

export default App;
