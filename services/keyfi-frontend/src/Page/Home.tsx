import './Home.css'


export default function Home() {


  return(
    <>
      <div className='home-page-container'>
        <h1 className='home-title'>Real estate made <span className='gradient-text'>simple</span></h1>
        <h3 className='home-description'>Buy, sell, tour virtually, get approved <br /> for loans: all powered by AI </h3>
        <input type='search' placeholder="Find your newest property" />
      </div>
    </>
  );
}