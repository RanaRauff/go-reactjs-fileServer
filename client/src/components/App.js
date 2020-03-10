import React from 'react';
// import './App.css';
import {Header} from "./Header"

import Actions from "./Actions"

class App extends React.Component{
  constructor(){
    super();
  }

  render(){
    return(
      <div>
        <Header/>
        <Actions/>
      </div>
    )
  }

}


// function App() {
//   return (
//     <div >
//       <Header/>
//     </div>
//   );
// }

export default App;
