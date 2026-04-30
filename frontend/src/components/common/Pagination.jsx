export default function Pagination({ page, totalPage, onChange }) {
  return (
    <div>
      <button disabled={page === 1} onClick={() => onChange(page - 1)}>
        Prev
      </button>

      <span>
        {page} / {totalPage || 1}
      </span>

      <button
        disabled={page >= totalPage}
        onClick={() => onChange(page + 1)}
      >
        Next
      </button>
    </div>
  );
}