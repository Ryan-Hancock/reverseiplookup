import React from "react";
import { Container, Row, Col } from "reactstrap";
import SearchBar from "./components/SearchBar.js";
import DomainTable from "./components/DomainTable.js";
import ApiHandler from "./apiHandler.js";

function App() {
  const handleSearchClick = (e) => {
    console.log(e);
    ApiHandler.searchByDomain(e);
  };

  return (
    <Container>
      <Row>
        <Col>
          <h2>Reverse IP Lookup</h2>
        </Col>
      </Row>
      <Row>
        <Col>
          <SearchBar onClick={handleSearchClick}></SearchBar>
          <DomainTable
            rows={[{ domain: "test", validFrom: "2020-05-20" }]}
          ></DomainTable>
        </Col>
      </Row>
    </Container>
  );
}

export default App;
