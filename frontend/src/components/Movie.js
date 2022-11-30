import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";

const Movie = () => {
  const [movie, setMovie] = useState({});
  let { id } = useParams();

  useEffect(() => {
    let myMovies = {
      id: 1,
      title: "Highlander",
      release_date: "1986-03-07",
      mpaa_rating: "R",
      description: "Some long description",
    };

    setMovie(myMovies);
  }, [id]);

  return (
    <div>
      <h2>
        Movie title: <strong>{movie.title}</strong>
      </h2>
      <hr />
    </div>
  );
};

export default Movie;
