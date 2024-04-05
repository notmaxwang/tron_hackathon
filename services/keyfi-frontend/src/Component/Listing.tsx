import './Listing.css';

export default function Listing(props: any) {
  const listingName = props.name;
  return(
    <>
      <div className="listing-container">
        <p>{listingName}</p>
      </div>
    </>
  );
}