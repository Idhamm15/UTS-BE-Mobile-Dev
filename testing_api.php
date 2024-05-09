<?php

// Endpoint yang ingin diakses
$endpoint = 'http://localhost:8080/api/products';

// Lakukan permintaan HTTP ke endpoint
$response = file_get_contents($endpoint);

// Periksa apakah permintaan berhasil
if ($response === false) {
    echo 'Gagal melakukan permintaan ke endpoint.';
} else {
    // Ubah respons JSON menjadi array PHP
    $data = json_decode($response, true);

    // Periksa apakah penguraian JSON berhasil
    if ($data === null) {
        echo 'Gagal mengurai respons JSON.';
    } else {
        // Tampilkan data produk
        echo '<h1>Product List</h1>';
        echo '<table border="1">';
        echo '<tr><th>ID</th><th>Name</th><th>Price</th><th>Description</th></tr>';
        foreach ($data as $product) {
            echo '<tr>';
            echo '<td>' . $product['id'] . '</td>';
            echo '<td>' . $product['product_name'] . '</td>';
            echo '<td>' . $product['price'] . '</td>';
            echo '<td>' . $product['description'] . '</td>';
            echo '</tr>';
        }
        echo '</table>';
    }
}

?>
