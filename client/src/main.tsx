import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";

import App from "./App.tsx";
import Home from "./components/Home.tsx";
import Movies from "./components/Movies.tsx";
import ErrorPage from "./components/ErrorPage.tsx";
import Movie from "./components/Movie.tsx";
import Login from "./components/Login.tsx";
import Graphql from "./components/Graphql.tsx";
import Catalogue from "./components/Catalogue.tsx";
import EditMovie from "./components/EditMovie.tsx";
import Genres from "./components/Genres.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <Home /> },
      { path: "/movies", element: <Movies /> },
      { path: "/movies/:id", element: <Movie /> },
      { path: "/genres", element: <Genres /> },
      { path: "/admin/movie/0", element: <EditMovie /> },
      { path: "/catalogue", element: <Catalogue /> },
      { path: "/graphql", element: <Graphql /> },
      { path: "/login", element: <Login /> },
    ],
  },
]);

const root = ReactDOM.createRoot(document.getElementById("root")!);
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
