import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const Movies = () => {
  const [movies, setMovies] = useState([]);

  useEffect(() => {
    let movieList = [
      {
        id: 1,
        title: "Highlander",
        release_date: "1986-03-07",
        mpaa_rating: "R",
        description: "Some long description",
      },
      {
        id: 2,
        title: "Raiders of the Lost Ark",
        release_date: "1981-06-12",
        mpaa_rating: "PG-13",
        description: "Some long description",
      },
    ];

    setMovies(movieList);
  }, []);

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
