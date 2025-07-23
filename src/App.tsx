import './App.css'
import { createBrowserRouter, RouterProvider } from 'react-router'
import Footer from './components/footer/Footer'
import Nav from './components/nav/Nav'
import { ThemeProvider } from './contexts/ThemeContext'

export default function App() {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <div>Home Page</div>,
    },
    {
      path: '/about',
      element: <div>About Page</div>,
    }
  ])

  return (
    <ThemeProvider>
      <header>
        <Nav />
      </header>

      <main>
        <RouterProvider router={router} />
      </main>

      <footer>
        <Footer />
      </footer>
    </ThemeProvider>
  )
}
