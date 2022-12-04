import React, { useState, useEffect, useCallback } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

const Movies = () => {
  const [movies, setMovies] = useState([]);

  const fetchMoviesHandler = useCallback(async () => {
    try {
      let url = "http://localhost:8080/movies";
      const res = await axios.get(url, {
        headers: { Accept: "application/json" },
        method: "GET",
      });
      if (res.statusText !== "OK") {
        throw new Error("Something went wrong during fetching data");
      }

      const { data } = res;
      console.log(data);
      setMovies(data.movies);
    } catch (error) {
      console.log(error);
    }
  }, []);

  useEffect(() => {
    fetchMoviesHandler();
  }, [fetchMoviesHandler]);

  return (
    <div>
      <h2>Movie</h2>
      <hr />
      <table className="table table-striped table-hover">
        <thead>
          <tr>
            <th>Movie</th>
            <th>Release date</th>
            <th>Rating</th>
          </tr>
        </thead>
        <tbody>
          {movies.map((movie) => (
            <tr key={movie.id}>
              <td>
                <Link to={`/movies/${movie.id}`}>{movie.title}</Link>
              </td>
              <td>{movie.release_date}</td>
              <td>{movie.mpaa_rating}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Movies;
