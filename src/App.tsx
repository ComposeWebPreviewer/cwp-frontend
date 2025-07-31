import './App.css'
import { createBrowserRouter, RouterProvider } from 'react-router'
import Footer from './components/footer/Footer'
import Nav from './components/nav/Nav'
import { ThemeProvider } from './contexts/ThemeContext'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import View from './screens/view/View'

export const API_BASE_URL = 'https://20s5mqesgj.execute-api.ap-south-1.amazonaws.com/Prod';

const queryClient = new QueryClient()

const router = createBrowserRouter([
  {
    path: '/',
    element: <div>Home Page</div>,
  },
  {
    path: '/view/:id',
    Component: View
  },
  {
    path: '/404',
    element: <div>404 Not Found</div>
  }
])

export default function App() {
  return (
    <ThemeProvider>
      <header>
        <Nav />
      </header>

      <main>
        <QueryClientProvider client={queryClient}>
          <RouterProvider router={router} />
        </QueryClientProvider>
      </main>

      <footer>
        <Footer />
      </footer>
    </ThemeProvider>
  )
}
