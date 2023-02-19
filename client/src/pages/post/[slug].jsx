import Head from 'next/head'
import styles from '@/styles/Home.module.css'

export default function PostPage({ postData }) {

  return (
    <>
      <Head>
        <title>Blog in React and Next.js </title>
        <meta name="description" content="Blog website built with Nextjs" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      {postData.map((post) => (
        <div className={styles.container} key={post.id}>
          <h1>{post.title}</h1>
          <br />
          <p style={{ color: 'grey' }}>Author: {post.username}</p>
          <br />
          <p>{post.content}</p>
          <br />
          <p style={{ color: 'grey', textAlign: 'right' }}>Post created: {post.created}</p>
        </div>
      ))}


    </>
  )
}

export async function getServerSideProps({ query }) {
  const postTitle = (query.slug).replaceAll("-", " ")

  let data = {
    title: postTitle
  }

  const req = await fetch(`http://localhost:8080/post`, { method: "POST", body: JSON.stringify(data) });
  const incomingData = await req.json();

  return {
    props: { postData: incomingData },
  }
}