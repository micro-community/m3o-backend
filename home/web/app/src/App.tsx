import React from 'react';
import HeaderImg from './assets/header.png';
import StoreIcon from './assets/store-icon.png';
import MessagingIcon from './assets/messaging-icon.png';
import DebuggingIcon from './assets/debugging-icon.png';
import CI from './assets/CI.svg';
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
            <h1>The development platform<br />for <span className='underline'>cloud services</span></h1>
            <p>Unleash developer volocity. Leaverage the scalable infrastructure provided by M3O and empower your team to focus on product, not the infrastructure running it.</p>

            <div className='actions'>
              <button className='primary'>Get started</button>
              <button className='secondary'>Read the docs</button>
            </div>
          </div>

          <img src={HeaderImg} alt='' />
        </div>

        <div className='runtime'>
          <div className='inner'>
            <h2>Explore the Micro runtime</h2>

            <div className='cards'>
              <div className='card'>
                <img src={StoreIcon} alt='Store' />
                <h5>Store</h5>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. <br /><br />Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
              </div>

              <div className='card'>
                <img src={MessagingIcon} alt='Messaging' />
                <h5>Messaging</h5>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. <br /><br />Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
              </div>

              <div className='card'>
                <img src={DebuggingIcon} alt='Debugging' />
                <h5>Debugging</h5>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. <br /><br />Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
              </div>
            </div>
          </div>
        </div>

        <div className='github-actions'>
          <div className='left'>
            <i>New!</i>
            <h2>Continuous Integration with<br /> GitHub Actions</h2>
            <p>Automatically deploy services each time you push to GitHub with the micro/action. Running a staging enviroment in M3O? Configure the environment to use and everytime you push a commit to your branch, M3O will build and deploy your enviroment.</p>
            <button className='secondary'>Setup GitHub actions</button>
          </div>

          <img src={CI} alt='GitHub Actions' />
        </div>
      </div>
    );
  }
}