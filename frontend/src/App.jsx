import { useState, useEffect } from 'react';

function App() {
  const [status, setStatus] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchStatus = () => {
      fetch('http://localhost:8080/status')
        .then((res) => res.json())
        .then((data) => {
          setStatus(data);
          setError(null);
        })
        .catch(() => setError('Cannot reach server'));
    };

    fetchStatus();
    const interval = setInterval(fetchStatus, 1000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div style={{ fontFamily: 'sans-serif', padding: '2rem' }}>
      <h1>Cluster Status</h1>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {status && (
        <div style={{ border: '1px solid #ccc', padding: '1rem', borderRadius: '8px', maxWidth: '300px' }}>
          <p><strong>ID:</strong> {status.id}</p>
          <p><strong>Role:</strong> {status.role}</p>
          <p><strong>Term:</strong> {status.term}</p>
        </div>
      )}
    </div>
  );
}

export default App;