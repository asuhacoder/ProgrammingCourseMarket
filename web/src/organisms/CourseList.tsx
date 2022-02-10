import React from 'react';
import {
  Card,
  CardContent,
  Typography,
} from '@mui/material';

function CourseList() {
  return (
    <Card>
      <CardContent>
        <Typography variant="h4" component="div">
          Rust Get Started
        </Typography>
        <Typography variant="body2">
          <div>
            Rust is a multi-paradigm, general-purpose programming language designed for performance and safety, especially safe concurrency.[12][13] Rust is syntactically similar to C++,[14] but can guarantee memory safety by using a borrow checker to validate references.[15] Rust achieves memory safety without garbage collection, and reference counting is optional.[16][17] Rust has been called a systems programming language and in addition to high-level features such as functional programming it also offers mechanisms for low-level memory management.
            First appearing in 2010, Rust was designed by Graydon Hoare at Mozilla Research,[18] with contributions from Dave Herman, Brendan Eich, and others.[19][20] The designers refined the language while writing the Servo experimental browser engine[21] and the Rust compiler. Rust&apos;s major influences include C++, OCaml, Haskell, and Erlang.[5] It has gained increasing use and investment in industry, by companies including Amazon, Discord, Dropbox, Facebook (Meta), Google (Alphabet), and Microsoft.
            Rust has been voted the &quot;most loved programming language&quot; in the Stack Overflow Developer Survey every year since 2016, and was used by 7% of the respondents in 2021.[22]
          </div>
        </Typography>
      </CardContent>
    </Card>
  );
}

export default CourseList;
