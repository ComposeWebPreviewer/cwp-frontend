import './Nav.css'
import logo from '../../assets/logo.svg'
import ThemeToggle from '../theme/ThemeToggle';

export default function Nav() {
  return (
    <>
      <nav>
        <div>
          <img src={logo} alt='Logo' />
          <h1>Compose Web</h1>
        </div>

        <ThemeToggle />
      </nav>
      <hr />
    </>
  );
}
