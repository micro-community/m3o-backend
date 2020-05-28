import React from 'react';
import HeaderImg from './assets/header.png';
import './App.scss';

export default class App extends React.Component {
  render(): JSX.Element {
    return (
      <div className="App">
        <header>
          <h3 className='logo'>M3O</h3>

          <nav>
            <a href='/about'>About</a>
            <a href='/pricing'>Pricing</a>
            <a href='/docs'>Docs</a>
            <a href='/login' className='login'>Login</a>
          </nav>
        </header>

        <div className='hero'>
          <div className='left'>
            <h1>The development platform<br /> for <span className='underline'>cloud services</span></h1>
            <p>Unleash developer volocity. Leaverage the scalable infrastructure provided by M3O and empower your team to focus on product, not the infrastructure running it.</p>

            <div className='actions'>
              <button className='primary'>Get started</button>
              <button className='secondary'>Read the docs</button>
            </div>
          </div>

          <img src={HeaderImg} alt='' />
        </div>
      </div>
    );
  }
}