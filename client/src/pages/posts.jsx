
import { useEffect, useState } from 'react'
export default function Blog() {


  const[reactData, setReactData] = useState([]);
  useEffect(() => {
    fetch('http://localhost:8080/')
    .then(res => res.json())
    .then(data => {
      setReactData(data);
    }).catch((e) => {console.log(e)});
  }, []);

  return (
    <ul>
     {reactData.map((post) => (
          <tr>
            <td>{post.id}</td>
            <td>{post.username}</td>
            <td>{post.title}</td>
          </tr>
        ))}
     
    </ul>
  )
}

