import React from 'react';
import logo from './logo.svg';
import { MenuNav } from './sections/layout/menu';
import { TopNaviBar } from './sections/layout/header';
import { Container } from './sections/layout/main';
import { Provider } from 'mobx-react';
import { rootStore as stores} from './stores/rootStore';


function App() {
  return (
    <Provider {...stores}>
      <div style={{overflow:'hidden',height:'100%'}}>
        <TopNaviBar/>
        <div style={{display:"flex",flexDirection: "row",height:'100%'}}>
          <MenuNav/>
          <Container style={{margin:'0 20px 0 20px',width:'100%',overflow:'auto'}}/>
        </div>
      </div>
    </Provider>
  );
}

export default App;
