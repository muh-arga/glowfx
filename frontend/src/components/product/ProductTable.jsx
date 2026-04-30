export default function ProductTable({ data, onDelete, onEdit }) {
  return (
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>SKU</th>
          <th>Status</th>
          <th>Action</th>
        </tr>
      </thead>

      <tbody>
        {data.length === 0 && (
          <tr>
            <td colSpan="5">No data</td>
          </tr>
        )}

        {data.map((p) => (
          <tr key={p.id}>
            <td>{p.id}</td>
            <td>{p.name}</td>
            <td>{p.sku}</td>
            <td>{p.status}</td>
            <td className="table-actions">
              <button className="secondary" onClick={() => onEdit(p)}>
                Edit
              </button>
              <button className="danger" onClick={() => onDelete(p.id)}>
                Delete
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}