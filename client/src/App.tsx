import "./App.css";
import { Outlet } from "react-router-dom";

function App() {
  return (
    <>
      <div>
        {/* <h2>Click on the Vite and React logos to learn more</h2> */}
        <Outlet />
      </div>
    </>
  );
}

export default App;
