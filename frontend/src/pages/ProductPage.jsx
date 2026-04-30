import { useState } from "react";
import useProducts from "../hooks/useProduct";
import { deleteProduct } from "../api/productApi";

import SearchInput from "../components/common/SearchInput";
import Pagination from "../components/common/Pagination";
import ProductTable from "../components/product/ProductTable";
import ProductFormModal from "../components/product/ProductFormModal";

export default function ProductPage() {
  const [search, setSearch] = useState("");
  const [page, setPage] = useState(1);
  const limit = 10;

  const [modalOpen, setModalOpen] = useState(false);
  const [selected, setSelected] = useState(null);

  const { products, total, loading, refetch } = useProducts(
    search,
    page,
    limit,
  );

  const totalPage = Math.ceil(total / limit);

  const handleDelete = async (id) => {
    if (!confirm("Delete this product?")) return;
    await deleteProduct(id);
    refetch();
  };

  const handleCreate = () => {
    setSelected(null);
    setModalOpen(true);
  };

  const handleEdit = (product) => {
    setSelected(product);
    setModalOpen(true);
  };

  return (
    <div className="container">
      <div className="header">
        <h2>Product</h2>
        <button className="primary" onClick={handleCreate}>
          + Create
        </button>
      </div>

      <div style={{ marginTop: 10 }}>
        <input
          placeholder="Search product..."
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
        />
      </div>

      {loading ? (
        <p>Loading...</p>
      ) : (
        <ProductTable
          data={products}
          onDelete={handleDelete}
          onEdit={handleEdit}
        />
      )}

      <Pagination page={page} totalPage={totalPage} onChange={setPage} />

      <ProductFormModal
        isOpen={modalOpen}
        onClose={() => setModalOpen(false)}
        onSuccess={refetch}
        initialData={selected}
      />
    </div>
  );
}
