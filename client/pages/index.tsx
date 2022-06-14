import type { NextPage } from 'next'
import { useState } from 'react'
import {postUpload, getViewAll} from '../apis/callbacks'
import styles from '../styles/Home.module.css'

const Home: NextPage = () => {
  
  const [fileSaved, setFileSaved] = useState<File | null>(null)
  const [labelMsg, setLabelMsg] = useState<string>('')

  const handleChange = (event: { target: { files: any } }) => {
    const { files } = event.target
    if (files && files.length > 0) {
      setFileSaved(files[0])
    }
  }

  const handleSubmit = async(event: { preventDefault: () => void }) => {
    event.preventDefault()

    const formData = new FormData()
    formData.append('file', fileSaved !== null ? fileSaved : new Blob())
    const response = await postUpload(formData)
    setLabelMsg(JSON.stringify(response))
    console.log(response)
  }

  const [viewAll, setViewAll] = useState<string[]>(["None"])
  const handleViewAll = async() => {
    const response = await getViewAll()
    setViewAll(response['files'])
  }

  return (
    <div className={styles.pagecenter}>
      <h2>Upload</h2>
      <label className='txt'>{labelMsg}</label>
      <form className={styles.form} onSubmit={handleSubmit}>
        <input type="file" name="file" className={styles.input} onChange={handleChange}/>
        <button type="submit">Upload</button>
      </form>
      <hr />
      <div className={styles.container}>
        <h2>View All</h2>
        <button onClick={handleViewAll}>View All</button>
      </div>
        <ul>
          {viewAll.map((item, index) => {
            return <li key={index}><a href={"http://localhost:8080/api/view/" + item}>{item.split('__').slice(1)}</a></li>
          }
          )}
        </ul>  
    </div>
  )
  
}


export default Home
