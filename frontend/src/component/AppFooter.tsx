import { Footer } from 'react-daisyui'

const AppFooter = () => {
  return (
    <footer>
      <Footer className="items-center justify-center p-4">
        <aside className="grid-flow-col items-center">
          <span className="text-lg font-bold">GitPulse</span>
          <p>Copyright © {new Date().getFullYear()} - All right reserved.</p>
          {import.meta.env.VITE_ENABLE_MOCK && <p>Network Mock enabled.</p>}
        </aside>
      </Footer>
    </footer>
  )
}

export default AppFooter
