import './App.css'
import Footer from './components/footer/Footer'
import Nav from './components/nav/Nav'
import { ThemeProvider } from './contexts/ThemeContext'

export default function App() {
  return (
    <ThemeProvider>
      <header>
        <Nav />
      </header>

      <Footer />
    </ThemeProvider>
  )
}
