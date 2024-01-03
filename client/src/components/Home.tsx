import { Link } from "react-router-dom";

const Home = () => {
  return (
    <div className="text-center">
      <h1>This is the home page</h1>
      <Link rel="stylesheet" to="/movies"></Link>
    </div>
  );
};

export default Home;
