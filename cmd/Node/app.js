const express = require('express')
const app = express()
const port = 3000

const { Pool } = require("pg");

// Create a PostgreSQL connection pool at startup
const pool = new Pool({
  host: "localhost",
  port: 5432,
  user: "postgres",
  password: "root",
  database: "postgres",
  max: 10,
  idleTimeoutMillis: 30000, // Close idle connections after 30s
});

// Check database connection at startup
pool.connect()
  .then(client => {
    console.log("Connected to PostgreSQL");
    client.release();
  })
  .catch(err => console.error("Failed to connect to PostgreSQL", err));

// API Route to get PostgreSQL version
app.get("/postgres", async (req, res) => {
  
});

function testEven( x ) {
    return (x % 2 == 0)
}

app.get('/js-health', (req, res) => {
  res.send("Running");
})

app.get('/js-rand', (req, res) => {
    let ab = Math.random(10)
    res.send({"number": ab});
})

app.get('/js-test-even/:num', (req, res) => {
    res.send({"is_even": testEven(req.params.num)});
})

app.get('/js-db', async (req, res) => {
  try {
    const result = await pool.query("SELECT version();");
    res.json({ version: result.rows[0].version });
  } catch (err) {
    res.status(500).json({ error: "Failed to fetch PostgreSQL version", details: err.message });
  }
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})


// const DemoGoRoutines = async (howmany) => {
//     const promises = [];
  
//     for (let i = 0; i < howmany; i++) {
//       promises.push(
//         new Promise((resolve) => {
//           setTimeout(() => {
//             console.log(`Running routine ${i}`);
//             resolve();
//           }, 5000); // simulate a 5-second delay
//         })
//       );
//     }
  
//     await Promise.all(promises); // Wait for all routines to finish
//   };
  
//   const main = async () => {
//     const start = Date.now();
//     console.log(`Starting my program at ${new Date(start).toISOString()}`);
  
//     const howmany = 100000;
//     await DemoGoRoutines(howmany);
  
//     const end = Date.now();
//     console.log(`Ended my program after ${((end - start) / 1000).toFixed(2)} seconds`);
//   };
  
//   main();